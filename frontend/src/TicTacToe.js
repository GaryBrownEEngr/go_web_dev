// https://reactjs.org/tutorial/tutorial.html

import React, { useState } from "react";
import { Link } from "react-router-dom";

import "./index.css";

function Square(props) {
  return (
    <button
      style={{
        backgroundColor: props.isWinningCell ? "Green" : "inherit",
      }}
      className="square"
      onClick={props.onClick}
    >
      {props.value}
    </button>
  );
}

function Board(props) {
  const out = [];
  for (let row = 0; row < 3; row++) {
    let inner = [];

    for (let col = 0; col < 3; col++) {
      let cell = col + 3 * row;
      inner.push(
        <Square
          key={cell}
          value={props.squares[cell]}
          isWinningCell={-1 !== props.partsOfWin.indexOf(cell)}
          onClick={() => props.onClick(cell)}
        />
      );
    }
    out.push(
      <div key={row} className="board-row">
        {inner}
      </div>
    );
  }

  return <div>{out}</div>;
}

function Game() {
  const [history, setHistory] = useState([
    {
      squares: Array(9).fill(null),
      selectedCellNumber: null,
    },
  ]);

  const [stepNumber, setStepNumber] = useState(0);
  const [xIsNext, setXIsNext] = useState(true);
  const [sortReversed, setSortReversed] = useState(false);

  function handleClick(i) {
    const historyNew = history.slice(0, stepNumber + 1);
    const current = historyNew[historyNew.length - 1];
    const squares = current.squares.slice();
    if (calculateWinner(squares) || squares[i]) {
      return;
    }
    squares[i] = xIsNext ? "X" : "O";

    setHistory(
      historyNew.concat([
        {
          squares: squares,
          selectedCellNumber: i,
        },
      ])
    );

    setStepNumber(historyNew.length);
    setXIsNext(!xIsNext);
  }

  function jumpTo(step) {
    setStepNumber(step);
    setXIsNext(step % 2 === 0);
  }

  const current = history[stepNumber];
  const winner = calculateWinner(current.squares);
  const partsOfWin = returnWinLine(current.squares);
  const isDraw = !winner && stepNumber === 9;

  const moves = history.map((step, move) => {
    const desc = move
      ? "Go to move #" +
        move +
        " r/c=" +
        (Math.floor(step.selectedCellNumber / 3) + 1) +
        "/" +
        ((step.selectedCellNumber % 3) + 1)
      : "Go to game start";
    return (
      <li key={move}>
        <button
          style={{
            fontWeight: move === stepNumber ? "bold" : "normal",
          }}
          onClick={() => jumpTo(move)}
        >
          {desc}
        </button>
      </li>
    );
  });
  if (sortReversed) {
    moves.reverse();
  }

  let status;
  if (winner) {
    status = "Winner: " + winner;
  } else if (isDraw) {
    status = "Draw";
  } else {
    status = "Next player: " + (xIsNext ? "X" : "O");
  }
  return (
    <React.Fragment>
      <Link to="/">Homepage</Link>
      <div className="game">
        <div className="game-board">
          <Board
            squares={current.squares}
            partsOfWin={partsOfWin}
            onClick={(i) => handleClick(i)}
          />
        </div>
        <div className="game-info">
          <div>{status}</div>
          <div>
            <label>
              reverse sort
              <input
                type="checkbox"
                checked={sortReversed}
                onChange={() => setSortReversed(!sortReversed)}
              ></input>
            </label>
          </div>

          <ol>{moves}</ol>
        </div>
      </div>
    </React.Fragment>
  );
}

function calculateWinner(squares) {
  const lines = [
    [0, 1, 2],
    [3, 4, 5],
    [6, 7, 8],
    [0, 3, 6],
    [1, 4, 7],
    [2, 5, 8],
    [0, 4, 8],
    [2, 4, 6],
  ];
  for (let i = 0; i < lines.length; i++) {
    const [a, b, c] = lines[i];
    if (squares[a] && squares[a] === squares[b] && squares[a] === squares[c]) {
      return squares[a];
    }
  }
  return null;
}

function returnWinLine(squares) {
  const lines = [
    [0, 1, 2],
    [3, 4, 5],
    [6, 7, 8],
    [0, 3, 6],
    [1, 4, 7],
    [2, 5, 8],
    [0, 4, 8],
    [2, 4, 6],
  ];
  for (let i = 0; i < lines.length; i++) {
    const [a, b, c] = lines[i];
    if (squares[a] && squares[a] === squares[b] && squares[a] === squares[c]) {
      return lines[i];
    }
  }
  return [];
}

export default Game;
