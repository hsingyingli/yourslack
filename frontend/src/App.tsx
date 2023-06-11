import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';
import HomePage from './pages/HomePage';
import SignInPage from './pages/SignInPage';
import {RequireAuth, RequireNotAuth} from './middlewares'
import SignUpPage from './pages/SignUpPage';

function App() {
  return (
    <Router>
      <Routes>
        <Route element={<RequireNotAuth/>}>
          <Route path="/signin" element={<SignInPage />} />
          <Route path="/signup" element={<SignUpPage/>} />
        </Route>
        {/*   Protect Routes   */}
        <Route element={<RequireAuth />}>
          <Route path="/" element={<HomePage />} />
        </Route>
        {/*   Protect Routes   */}
      </Routes>
    </Router>
  );
}

export default App;
