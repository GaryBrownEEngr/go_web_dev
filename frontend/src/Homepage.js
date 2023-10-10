import React from "react";
import { Link } from "react-router-dom";

export default function Homepage() {
  return (
    <div>
      <h1>Homepage!</h1>
      <ul>
        <li>
          <Link to="/articles">Articles</Link>
        </li>
        <li>
          <Link to="/tictactoe">TicTacToe</Link>
        </li>
      </ul>
    </div>
  );
}
