import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import ChakraProvider from './providers/ChakraProvider';
import MainLayout from './components/Layouts/MainLayout';
import AuthProvider from './providers/AuthProvider';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

root.render(
  <React.StrictMode>
    <ChakraProvider>
      <AuthProvider>
        <MainLayout>
          <App />
        </MainLayout>
      </AuthProvider>
    </ChakraProvider>
  </React.StrictMode>
);

