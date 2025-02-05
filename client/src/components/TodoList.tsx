import { Flex, Spinner, Stack, Text } from "@chakra-ui/react";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";
import { BASE_URL } from "../App";
import TodoItem from "./TodoItem";

export type Todo = {
  id: string;
  userId: string;
  title: string;
  complete: boolean;
};

const TodoList = () => {
  const token = localStorage.getItem("token"); // Retrieve the token from localStorage

  const { data: todos, isLoading, isError } = useQuery<Todo[]>({
    queryKey: ["todo"],
    queryFn: async () => {
      try {
        const res = await axios.get(BASE_URL + "/gettodos", {
          headers: {
            "Authorization": `Bearer ${token}`, // Add the token to the request header
          },
        });

        return res.data.todo || [];
      } catch (error: unknown) {
        if (axios.isAxiosError(error)) {
          // Handle Axios error here
          throw new Error(error.response?.data?.error || "Something went wrong");
        }
        // Handle unknown errors
        throw new Error("An unexpected error occurred");
      }
    },
  });

  // Handle unauthenticated or unauthorized access
  if (isError) {
    return (
      <Text color="red.500" textAlign="center">
        Unauthorized. Please log in to access your tasks.
      </Text>
    );
  }

  return (
    <>
      <Text
        fontSize={"4xl"}
        textTransform={"uppercase"}
        fontWeight={"bold"}
        textAlign={"center"}
        my={2}
        bgGradient="linear(to-l, #0b85f8, #00ffff)"
        bgClip="text"
      >
        Today's Tasks
      </Text>
      {isLoading && (
        <Flex justifyContent={"center"} my={4}>
          <Spinner size={"xl"} />
        </Flex>
      )}
      {!isLoading && todos?.every(todo => todo.complete) && (
        <Stack alignItems={"center"} gap="3">
          <Text fontSize={"xl"} textAlign={"center"} color={"gray.500"}>
            All tasks completed! 🤞
          </Text>
          <img src="/go.png" alt="Go logo" width={70} height={70} />
        </Stack>
      )}
      <Stack gap={3}>
        {todos?.map((todo) => (
          <TodoItem key={todo.id} todo={todo} />
        ))}
      </Stack>
    </>
  );
};

export default TodoList;
