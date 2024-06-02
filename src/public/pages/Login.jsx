import { Box, TextField, Button, Typography } from "@mui/material";
import Form from "../components/Form";
import { LoginService } from "../../service/LoginService";

export default function Login() {
  const loginServices = new LoginService();
  return (
    <div className="row">
      {/* Esquerda */}
      <Box
        className="col-md-6 boxEsquerda"
        sx={{
          backgroundColor: "red",
          height: "100vh",
          alignItems: "center",
          justifyContent: "center",
        }}
      >
        <div className="logo"></div>
      </Box>

      {/* Direita */}
      <Box
        className="col-md-6 col-12"
        sx={{
          backgroundColor: "white",
          height: "100vh",
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
        }}
      >
        <Box
          sx={{
            border: "1px solid black",
            borderRadius: "20px",
            padding: "35px",
            boxShadow: "0.1em 0.1em 1em black",
          }}
        >
          <Form submitHandler={loginServices.Login}>
            <Typography
              sx={{ marginBottom: "5px" }}
              variant="h4"
              color="initial"
            >
              Login
            </Typography>
            <TextField
              sx={{ marginBottom: "5px" }}
              placeholder="Email"
              type="email"
              name="email"
              required
            />{" "}
            <br />
            <TextField
              sx={{ marginBottom: "5px" }}
              placeholder="Password"
              type="password"
              name="password"
              required
            />{" "}
            <br />
            <Button
              sx={{ float: "right" }}
              type="submit"
              variant="contained"
              color="primary"
            >
              Submit
            </Button>
          </Form>
        </Box>
      </Box>
    </div>
  );
}
