import React from 'react';
import {
  Box,
  Flex,
  Heading,
  Spacer,
  Button,
  useColorModeValue,
  useColorMode
} from '@chakra-ui/react';
import { MoonIcon, SunIcon } from '@chakra-ui/icons';
import { useNavigate } from 'react-router-dom';

const Header: React.FC = () => {
  const { colorMode, toggleColorMode } = useColorMode();
  const navigate = useNavigate();

  const bgColor = useColorModeValue('gray.100', 'gray.700');
  const textColor = useColorModeValue('gray.800', 'white');

  const handleLogout = () => {
    // Clear authentication token
    localStorage.removeItem('token');
    navigate('/login');
  };

  return (
    <Box 
      bg={bgColor} 
      py={4} 
      px={8} 
      color={textColor}
      boxShadow="md"
    >
      <Flex align="center" maxW="container.xl" mx="auto">
        <Heading size="lg" cursor="pointer" onClick={() => navigate('/')}>
          MyApp
        </Heading>
        <Spacer />
        <Flex align="center">
          <Button 
            onClick={toggleColorMode} 
            variant="ghost" 
            mr={4}
          >
            {colorMode === 'light' ? <MoonIcon /> : <SunIcon />}
          </Button>
          <Button 
            colorScheme="red" 
            variant="outline"
            onClick={handleLogout}
          >
            Logout
          </Button>
        </Flex>
      </Flex>
    </Box>
  );
};

export default Header;