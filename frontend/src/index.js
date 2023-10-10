import React, { useState } from "react";
import ReactDOM from "react-dom";
import { Route, Routes, BrowserRouter } from "react-router-dom";
import Homepage from "./Homepage";
import TicTacToe from "./TicTacToe";
import Articles from "./Articles";
import Article from "./Article";
import ArticleNew from "./ArticleNew";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Homepage />} />
        <Route path="/tictactoe" element={<TicTacToe />} />
        <Route path="/articles" element={<Articles />} />
        <Route path="/article/:articleId" element={<Article />} />
        <Route path="/article/new" element={<ArticleNew />} />
      </Routes>
    </BrowserRouter>
  );
}

// ========================================
ReactDOM.render(<App />, document.getElementById("root"));
