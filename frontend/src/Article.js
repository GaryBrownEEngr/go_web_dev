import React, { useState, useEffect } from "react";
import { Link, useParams, useNavigate } from "react-router-dom";

export default function Article() {
  const [articleData, setArticleData] = useState(undefined);
  const navigate = useNavigate();

  const articleId = useParams().articleId;
  useEffect(() => {
    (async function () {
      const response = await fetch(
        "http://localhost:10000/api/article/" + articleId,
        {
          method: "GET",
        }
      );
      const data = await response.json();
      setArticleData(data);
    })();
  }, [articleId]);

  async function DeleteArticle() {
    await fetch("http://localhost:10000/api/article/" + articleId, {
      method: "DELETE",
    });
    navigate("/articles");
  }

  return (
    <React.Fragment>
      <Link to="/">Homepage</Link>
      <button style={{ float: "right" }} onClick={DeleteArticle}>
        Delete
      </button>
      {articleData && (
        <div>
          <h1>{articleData.Title}</h1>
          <p>{articleData.content}</p>
        </div>
      )}
    </React.Fragment>
  );
}
