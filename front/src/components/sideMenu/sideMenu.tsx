// SideMenu.js
import { Link } from 'react-router-dom';
import { FaHome } from "react-icons/fa";
import { FaUser } from "react-icons/fa6";
import { FaStar } from "react-icons/fa";

const SideMenu = () => {
  return (
    <div className="h-full fixed bg-gray-100 border-r border-gray-300 w-52 flex flex-col p-5 top-0 left-0">
      <ul className="list-none p-0">
        <li className="my-2.5">
          <Link to="/" className="no-underline text-gray-800 text-lg p-2.5 flex items-center rounded transition-colors duration-300 hover:bg-gray-300">
            <FaHome className="mr-2" />Home
          </Link>
        </li>
        <li className="my-2.5">
          <Link to="/myPage" className="no-underline text-gray-800 text-lg p-2.5 flex items-center rounded transition-colors duration-300 hover:bg-gray-300">
            <FaUser className="mr-2" />My Page
          </Link>
        </li>
        <li className="my-2.5">
          <Link to="/favorite" className="no-underline text-gray-800 text-lg p-2.5 flex items-center rounded transition-colors duration-300 hover:bg-gray-300">
            <FaStar className="mr-2" />Favorite
          </Link>
        </li>
      </ul>
    </div>
  );
};

export default SideMenu;