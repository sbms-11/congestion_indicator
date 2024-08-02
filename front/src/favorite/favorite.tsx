import SideMenu from "../components/sideMenu/sideMenu"
import Table from "./table"

const Favorite = () => {
    return (
        <>
            <SideMenu />
            <div className="pl-[210px] h-full">
                <Table />
            </div>
        </>
    );
};

export default Favorite;