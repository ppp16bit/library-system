import React from 'react';
import BookCard from './BookCard';
import './BookList.css';

const BookList = ({ books, onEditBook, onDeleteBook }) => {
  if (!books) {
    return (
      <div className="loading-container">
        <p>Carregando livros...</p>
      </div>
    );
  }

  if (books.length === 0) return <p className="no-books-message">Nenhum livro cadastrado.</p>;

  return (
    <div className="book-list-container">
      {books.map(book => (
        <BookCard
          key={book.id}
          book={book}
          onEdit={() => onEditBook(book)}
          onDelete={() => onDeleteBook(book.id)}
        />
      ))}
    </div>
  );
};

export default BookList;
