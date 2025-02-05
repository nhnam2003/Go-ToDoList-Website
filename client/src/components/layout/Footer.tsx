import React from 'react';
import {
  Box,
  Container,
  Stack,
  Text,
  useColorModeValue
} from '@chakra-ui/react';

const Footer: React.FC = () => {
  const bgColor = useColorModeValue('gray.100', 'gray.700');
  const textColor = useColorModeValue('gray.600', 'gray.200');

  return (
    <Box
      bg={bgColor}
      color={textColor}
      py={6}
      position="absolute"
      bottom="0"
      width="full"
    >
      <Container 
        as={Stack} 
        maxW={'container.xl'} 
        direction={{ base: 'column', md: 'row' }}
        spacing={4}
        justify={{ base: 'center', md: 'space-between' }}
        align={{ base: 'center', md: 'center' }}
      >
        <Text>© 2024 MyApp. All rights reserved</Text>
        <Stack direction={'row'} spacing={6}>
          <Text>Privacy Policy</Text>
          <Text>Terms of Service</Text>
          <Text>Contact</Text>
        </Stack>
      </Container>
    </Box>
  );
};

export default Footer;