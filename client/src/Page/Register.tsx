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

const RegisterPage: React.FC = () => {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [error, setError] = useState("");

  const navigate = useNavigate();

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    // Basic validation
    if (password !== confirmPassword) {
      setError("Passwords do not match");
      return;
    }

    try {
      const response = await axios.post("/register", {
        username,
        email,
        password,
      });

      // Store token or handle successful registration
      localStorage.setItem("token", response.data.token);
      navigate("/dashboard");
    } catch (err) {
      setError("Registration failed. Please try again.");
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
          Register
        </Heading>
        <form onSubmit={handleRegister}>
          <VStack spacing={4}>
            <FormControl id="username" isRequired>
              <FormLabel>Username</FormLabel>
              <Input
                type="text"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                placeholder="Choose a username"
              />
            </FormControl>
            <FormControl id="email" isRequired>
              <FormLabel>Email</FormLabel>
              <Input
                type="text"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                placeholder="Enter your email"
              />
            </FormControl>
            <FormControl id="password" isRequired>
              <FormLabel>Password</FormLabel>
              <Input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="Create a password"
              />
            </FormControl>
            <FormControl id="confirm-password" isRequired>
              <FormLabel>Confirm Password</FormLabel>
              <Input
                type="password"
                value={confirmPassword}
                onChange={(e) => setConfirmPassword(e.target.value)}
                placeholder="Confirm your password"
              />
            </FormControl>
            {error && (
              <Text color="red.500" textAlign="center">
                {error}
              </Text>
            )}
            <Button colorScheme="blue" width="full" type="submit">
              Register
            </Button>
            <Text>
              Already have an account?{" "}
              <ChakraLink color="blue.500" onClick={() => navigate("/login")}>
                Log In
              </ChakraLink>
            </Text>
          </VStack>
        </form>
      </Box>
    </Box>
  );
};

export default RegisterPage;
