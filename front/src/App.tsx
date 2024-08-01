import { BrowserRouter as Router, Route, Routes} from "react-router-dom"
import './App.css'
import Home from "./home/home";
import MyPage from "./myPage/myPage";
import ShopDetail from "./shopDetail/shopDetail";

function App() {

  return (
    <>
      <Router>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/myPage" element={<MyPage />} />
          <Route path="/shopDetail" element={<ShopDetail />} />
        </Routes>
      </Router>
    </>
  )
}

export default App
