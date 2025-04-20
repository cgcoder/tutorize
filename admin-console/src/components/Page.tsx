import { Colors } from "@blueprintjs/core";

export interface PageProps {
    children: React.ReactNode;
}
export default function Page(props: PageProps) {
    return (
        <div style={{ display: "flex", width: "100%", background: Colors.WHITE, padding: "10px" }} data-type="page">
            {props.children}
        </div>
    );
}