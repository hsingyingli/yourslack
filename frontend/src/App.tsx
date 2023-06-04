import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';
import HomePage from './pages/Home';
import {RequireAuth} from './middlewares'

function App() {
  return (
    <Router>
      <Routes>
      {/*
        <Route exact path="/login" element={<SignIn />} />
        <Route exact path="/register" element={<Register />} />
      */}
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
