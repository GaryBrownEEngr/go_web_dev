import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

export default function ArticleNew() {
  const [title, setTitle] = useState("");
  const [desc, setDesc] = useState("");
  const [content, setContent] = useState("");
  const navigate = useNavigate();

  async function createArticle() {
    let request = { Title: title, desc: desc, content: content };
    const response = await fetch("http://localhost:10000/api/article", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(request),
    });
    const data = await response.json();
    navigate("/article/" + data.Id);
  }

  return (
    <React.Fragment>
      <Link to="/">Homepage</Link>
      <div>
        <h2>Title</h2>
        <input
          type="text"
          value={title}
          onChange={(v) => setTitle(v.target.value)}
        ></input>
        <h2>Discription</h2>
        <input
          type="text"
          value={desc}
          onChange={(v) => setDesc(v.target.value)}
        ></input>
        <h2>Content</h2>
        <textarea
          value={content}
          onChange={(v) => setContent(v.target.value)}
          rows="50"
          style={{ width: "100%" }}
        ></textarea>
      </div>
      <button onClick={createArticle}>Create</button>
    </React.Fragment>
  );
}
