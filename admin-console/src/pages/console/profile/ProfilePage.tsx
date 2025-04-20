import { Card, H3 } from "@blueprintjs/core";
import { Link, Outlet } from "react-router";

export default function ProfilePage() {
    return (
        <Card>
            <H3>Profile</H3>
            <p>Manage your profile information here.</p>

            <ul style={{ listStyle: "none", padding: 0 }}>
                <li>
                    <Link to="create">Create Profile</Link>
                </li>
                <li>
                    <Link to="edit">Edit Profile</Link>
                </li>
            </ul>
            <Outlet />
        </Card>
    );
}