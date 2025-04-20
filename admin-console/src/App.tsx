import "normalize.css";
import "@blueprintjs/core/lib/css/blueprint.css";
import "@blueprintjs/icons/lib/css/blueprint-icons.css";
import "./App.css";
import {
  Routes,
  Route,
  BrowserRouter,
} from "react-router";
import ConsolePage from "./pages/console/ConsolePage";
import LoginPage from "./pages/login/LoginPage";
import DashboardPage from "./pages/console/dashboard/DashboardPage";
import ProfilePage from "./pages/console/profile/ProfilePage";
import SettingsPage from "./pages/console/settings/SettingsPage";

function App() {
  return (
    <BrowserRouter>
        {/* Main Content Area */}
        <main style={{ flex: 1, margin: "0px" }}>
          <Routes>
            <Route path="/console" element={<ConsolePage />}>
              <Route index element={<DashboardPage />} />
              <Route path="profile" element={<ProfilePage />}>
                <Route path="create" element={<span>Create Profile</span>} />
                <Route path="edit" element={<span>Edit Profile</span>} />
              </Route>
              <Route path="settings" element={<SettingsPage />} />
            </Route>
            <Route path="/login" element={<LoginPage />} />
            <Route path="/register" element={<span>Register</span>} />
          </Routes>
        </main>
    </BrowserRouter>
  );
}

export default App;
