import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";

export default function Articles() {
  const [articlesData, setArticlesData] = useState(undefined);
  useEffect(() => {
    (async function () {
      const response = await fetch("http://localhost:10000/api/articles", {
        method: "GET",
      });
      const data = await response.json();
      setArticlesData(data);
    })();
  }, []);

  return (
    <React.Fragment>
      <Link to="/">Homepage</Link>
      <div>
        <h1>Article List</h1>
        {articlesData &&
          articlesData.map((item, index) => (
            <div key={index} style={{ borderTop: "solid" }}>
              <Link to={"/article/" + item.Id}>
                <h2>{item.Title}</h2>
                <p>{item.desc}</p>
              </Link>
            </div>
          ))}
      </div>
      <Link to="/article/new">
        <button>Create a New Article</button>
      </Link>
    </React.Fragment>
  );
}
