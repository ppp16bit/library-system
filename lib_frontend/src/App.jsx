import React, { useState, useEffect, useRef } from 'react';
import './App.css';
import BookList from './components/books/BookList';
import BookForm from './components/books/BookForm';
import UserList from './components/users/UserList';
import UserForm from './components/users/UserForm';
import LoanList from './components/loans/LoanList';
import LoanForm from './components/loans/LoanForm';
import axios from 'axios';
import { API_BASE_URL } from './constants';
import { CSSTransition } from 'react-transition-group';

function App() {
  const [currentView, setCurrentView] = useState('books');
  const [showForm, setShowForm] = useState(false);
  const [bookToEdit, setBookToEdit] = useState(null);
  const [userToEdit, setUserToEdit] = useState(null);
  const [loanToEdit, setLoanToEdit] = useState(null);
  const [books, setBooks] = useState([]);
  const [users, setUsers] = useState([]);
  const [loans, setLoans] = useState([]);
  const bookListDivRef = useRef(null);
  const bookFormRef = useRef(null);
  const userListDivRef = useRef(null);
  const userFormRef = useRef(null);
  const loanListDivRef = useRef(null);
  const loanFormRef = useRef(null);

  const fetchBooks = async () => {
    try {
      const response = await axios.get(`${API_BASE_URL}/books`, {
        validateStatus: (status) => status >= 200 && status < 300 || status === 204
      });
      setBooks(response.data);
    } catch (error) {
      console.error("Erro ao carregar livros:", error);
    }
  };

  const fetchUsers = async () => {
    try {
      const response = await axios.get(`${API_BASE_URL}/users`, {
        validateStatus: (status) => status >= 200 && status < 300 || status === 204
      });
      setUsers(response.data);
    } catch (error) {
      console.error("Erro ao carregar usuários:", error);
    }
  };

  const fetchLoans = async () => {
    try {
      const response = await axios.get(`${API_BASE_URL}/loans`, {
        validateStatus: (status) => status >= 200 && status < 300 || status === 204
      });
      setLoans(response.data);
    } catch (error) {
      console.error("Erro ao carregar empréstimos:", error);
    }
  };

  useEffect(() => {
    if (currentView === 'books') {
      fetchBooks();
    } else if (currentView === 'users') {
      fetchUsers();
    } else if (currentView === 'loans') {
      fetchLoans();
    }
  }, [currentView]);

  const handleSaveBook = async (bookData) => {
    try {
      if (bookData.id) {
        await axios.put(`${API_BASE_URL}/books/${bookData.id}`, bookData);
        alert('Livro atualizado com sucesso!');
      } else {
        await axios.post(`${API_BASE_URL}/books`, bookData);
        alert('Livro adicionado com sucesso!');
      }
      setShowForm(false);
      setBookToEdit(null);
      fetchBooks();
    } catch (error) {
      console.error("Erro ao salvar livro:", error.response ? error.response.data : error.message);
      alert('Erro ao salvar livro: ' + (error.response?.data?.message || 'Verifique o console para mais detalhes.'));
    }
  };

  const handleEditBook = (book) => {
    setBookToEdit(book);
    setShowForm(true);
  };

  const handleDeleteBook = async (id) => {
    if (window.confirm('Tem certeza que deseja deletar este livro?')) {
      try {
        await axios.delete(`${API_BASE_URL}/books/${id}`, {
          validateStatus: (status) => status === 200 || status === 204
        });
        alert('Livro deletado com sucesso!');
        fetchBooks();
      } catch (error) {
        console.error("Erro ao deletar livro:", error.response ? error.response.data : error.message);
        alert('Erro ao deletar livro: ' + (error.response?.data?.message || 'Verifique o console para mais detalhes.'));
      }
    }
  };

  const handleSaveUser = async (userData) => {
    try {
      if (userData.id) {
        await axios.put(`${API_BASE_URL}/users/${userData.id}`, userData);
        alert('Usuário atualizado com sucesso!');
      } else {
        await axios.post(`${API_BASE_URL}/users`, userData);
        alert('Usuário adicionado com sucesso!');
      }
      setShowForm(false);
      setUserToEdit(null);
      fetchUsers();
    } catch (error) {
      console.error("Erro ao salvar usuário:", error.response ? error.response.data : error.message);
      alert('Erro ao salvar usuário: ' + (error.response?.data?.message || 'Verifique o console para mais detalhes.'));
    }
  };

  const handleEditUser = (user) => {
    setUserToEdit(user);
    setShowForm(true);
  };

  const handleDeleteUser = async (id) => {
    if (window.confirm('Tem certeza que deseja deletar este usuário?')) {
      try {
        await axios.delete(`${API_BASE_URL}/users/${id}`, {
          validateStatus: (status) => status === 200 || status === 204
        });
        alert('Usuário deletado com sucesso!');
        fetchUsers();
      } catch (error) {
        console.error("Erro ao deletar usuário:", error.response ? error.response.data : error.message);
        alert('Erro ao deletar usuário: ' + (error.response?.data?.message || error.message || 'Verifique o console para mais detalhes.'));
      }
    }
  };

  const handleSaveLoan = async (loanData) => {
    try {
      if (loanData.id) {
        await axios.put(`${API_BASE_URL}/loans/${loanData.id}`, loanData);
        alert('Empréstimo atualizado com sucesso!');
      } else {
        await axios.post(`${API_BASE_URL}/loans`, loanData);
        alert('Empréstimo registrado com sucesso!');
      }
      setShowForm(false);
      setLoanToEdit(null);
      fetchLoans();
    } catch (error) {
      console.error("Erro ao salvar empréstimo:", error.response ? error.response.data : error.message);
      alert('Erro ao salvar empréstimo: ' + (error.response?.data?.message || 'Verifique o console para mais detalhes.'));
    }
  };

  const handleEditLoan = (loan) => {
    setLoanToEdit(loan);
    setShowForm(true);
  };

  const handleDeleteLoan = async (id) => {
    if (window.confirm('Tem certeza que deseja deletar este empréstimo?')) {
      try {
        await axios.delete(`${API_BASE_URL}/loans/${id}`, {
          validateStatus: (status) => status === 200 || status === 204
        });
        alert('Empréstimo deletado com sucesso!');
        fetchLoans();
      } catch (error) {
        console.error("Erro ao deletar empréstimo:", error.response ? error.response.data : error.message);
        alert('Erro ao deletar empréstimo: ' + (error.response?.data?.message || 'Verifique o console para mais detalhes.'));
      }
    }
  };

  const handleCancelForm = () => {
    setShowForm(false);
    setBookToEdit(null);
    setUserToEdit(null);
    setLoanToEdit(null);
    if (currentView === 'books') fetchBooks();
    if (currentView === 'users') fetchUsers();
    if (currentView === 'loans') fetchLoans();
  };

  const handleAddButton = () => {
    setBookToEdit(null);
    setUserToEdit(null);
    setLoanToEdit(null);
    setShowForm(true);
  };

  return (
    <div className="app-container">
      <h1>Sistema de Biblioteca</h1>

      <div className="navigation-buttons">
        <button
          onClick={() => { setCurrentView('books'); setShowForm(false); handleCancelForm(); }}
          className={currentView === 'books' ? 'active-nav-button' : 'nav-button'}
        >
          Gerenciar Livros
        </button>
        <button
          onClick={() => { setCurrentView('users'); setShowForm(false); handleCancelForm(); }}
          className={currentView === 'users' ? 'active-nav-button' : 'nav-button'}
        >
          Gerenciar Usuários
        </button>
        <button
          onClick={() => { setCurrentView('loans'); setShowForm(false); handleCancelForm(); }}
          className={currentView === 'loans' ? 'active-nav-button' : 'nav-button'}
        >
          Gerenciar Empréstimos
        </button>
      </div>

      {!showForm && currentView !== 'loans' && (
        <button onClick={handleAddButton} className="add-item-button">
          Adicionar Novo {currentView === 'books' ? 'Livro' : 'Usuário'}
        </button>
      )}
      {!showForm && currentView === 'loans' && (
        <button onClick={handleAddButton} className="add-item-button">
          Registrar Novo Empréstimo
        </button>
      )}

      <CSSTransition
        in={showForm && currentView === 'books'}
        timeout={300}
        classNames="fade"
        unmountOnExit
        nodeRef={bookFormRef}
      >
        <div ref={bookFormRef}>
          <BookForm
            bookToEdit={bookToEdit}
            onSubmit={handleSaveBook}
            onCancel={handleCancelForm}
          />
        </div>
      </CSSTransition>

      <CSSTransition
        in={showForm && currentView === 'users'}
        timeout={300}
        classNames="fade"
        unmountOnExit
        nodeRef={userFormRef}
      >
        <div ref={userFormRef}>
          <UserForm
            userToEdit={userToEdit}
            onSubmit={handleSaveUser}
            onCancel={handleCancelForm}
          />
        </div>
      </CSSTransition>

      <CSSTransition
        in={showForm && currentView === 'loans'}
        timeout={300}
        classNames="fade"
        unmountOnExit
        nodeRef={loanFormRef}
      >
        <div ref={loanFormRef}>
          <LoanForm
            loanToEdit={loanToEdit}
            onSubmit={handleSaveLoan}
            onCancel={handleCancelForm}
          />
        </div>
      </CSSTransition>

      <CSSTransition
        in={!showForm && currentView === 'books'}
        timeout={300}
        classNames="fade"
        unmountOnExit
        nodeRef={bookListDivRef}
      >
        <div ref={bookListDivRef}>
          <BookList
            books={books}
            onEditBook={handleEditBook}
            onDeleteBook={handleDeleteBook}
          />
        </div>
      </CSSTransition>

      <CSSTransition
        in={!showForm && currentView === 'users'}
        timeout={300}
        classNames="fade"
        unmountOnExit
        nodeRef={userListDivRef}
      >
        <div ref={userListDivRef}>
          <UserList
            users={users}
            onEditUser={handleEditUser}
            onDeleteUser={handleDeleteUser}
          />
        </div>
      </CSSTransition>

      <CSSTransition
        in={!showForm && currentView === 'loans'}
        timeout={300}
        classNames="fade"
        unmountOnExit
        nodeRef={loanListDivRef}
      >
        <div ref={loanListDivRef}>
          <LoanList
            loans={loans}
            onEditLoan={handleEditLoan}
            onDeleteLoan={handleDeleteLoan}
          />
        </div>
      </CSSTransition>
    </div>
  );
}

export default App;
