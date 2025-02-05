import React, { useState } from "react";
import {
  Box,
  Button,
  FormControl,
  FormLabel,
  Input,
  VStack,
  Heading,
  useColorModeValue,
  Text,
  Link as ChakraLink,
} from "@chakra-ui/react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

const LoginPage: React.FC = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    try {
      const response = await axios.post("/login", { username, password });
      if (response.data.token) {
        localStorage.setItem("token", response.data.token); // Lưu token vào LocalStorage
        return true;
      }
    } catch (error) {
      console.error("Lỗi khi đăng nhập", error);
      return false;
    }
  };

  const bgColor = useColorModeValue("gray.100", "gray.700");

  return (
    <Box
      minHeight="100vh"
      display="flex"
      alignItems="center"
      justifyContent="center"
      bg={bgColor}
    >
      <Box
        width="100%"
        maxWidth="400px"
        p={8}
        borderWidth={1}
        borderRadius="lg"
        boxShadow="lg"
        bg={useColorModeValue("white", "gray.800")}
      >
        <Heading textAlign="center" mb={6}>
          Login
        </Heading>
        <form onSubmit={handleLogin}>
          <VStack spacing={4}>
            <FormControl id="email" isRequired>
              <FormLabel>Email</FormLabel>
              <Input
                type="email"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                placeholder="Enter your email"
              />
            </FormControl>
            <FormControl id="password" isRequired>
              <FormLabel>Password</FormLabel>
              <Input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="Enter your password"
              />
            </FormControl>
            {error && (
              <Text color="red.500" textAlign="center">
                {error}
              </Text>
            )}
            <Button colorScheme="blue" width="full" type="submit">
              Log In
            </Button>
            <Text>
              Don't have an account?{" "}
              <ChakraLink
                color="blue.500"
                onClick={() => navigate("/register")}
              >
                Register
              </ChakraLink>
            </Text>
          </VStack>
        </form>
      </Box>
    </Box>
  );
};

export default LoginPage;
