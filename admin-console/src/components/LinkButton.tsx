import { Button, ButtonProps } from "@blueprintjs/core";
import { Link } from "react-router";
import { useLocation } from "react-router";

export interface LinkButtonProps {
  to: string;
  icon: ButtonProps["icon"];
  text: string;
  iconOnly?: boolean;
}

function routeMatches(currentPath: string, to: string): boolean {
  return to === "/console" ? currentPath === "/console" : currentPath.startsWith(to);
}

export default function LinkButton(props: LinkButtonProps) {
  const location = useLocation();
  return (
    <Link to={props.to}>
      <Button
        style={{
          paddingRight: props.iconOnly ? "10px" : "20px",
          margin: "0",
          padding: "10px",
          width: props.iconOnly ? "50px" : "150px",
        }}
        alignText="left"
        icon={props.icon}
        variant="solid"
        active={routeMatches(location.pathname, props.to)}
        text={props.iconOnly ? null : props.text}
        className="bp5-minimal"
      />
    </Link>
  );
}
