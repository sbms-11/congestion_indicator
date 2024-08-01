import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import './sideMenu.css';

const SideMenu: React.FC = () => {
    return (
      <nav className="w-64 h-screen bg-[#333] text-[#FFFFFF]">
        <ul>
          <li className="p-4 hover:bg-[#3C3C3C]">
            <a href="#home">HOME</a>
          </li>
          <li className="p-4 hover:bg-[#3C3C3C]">
            <a href="#mypage">MYPAGE</a>
          </li>
          <li className="p-4 hover:bg-[#3C3C3C]">
            <a href="#favorite">FAVORITE</a>
          </li>
        </ul>
      </nav>
    );
  };
  

export default SideMenu;