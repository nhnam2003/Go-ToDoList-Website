import { Badge, Box, Flex, Spinner, Text } from "@chakra-ui/react";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import { Todo } from "./TodoList";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { BASE_URL } from "../App";
import axios from "axios";

const TodoItem = ({ todo }: { todo: Todo }) => {
  const queryClient = useQueryClient();

  // Mutation để cập nhật trạng thái của todo
  const { mutate: updateTodo, isPending: isUpdating } = useMutation({
    mutationKey: ["updateTodo"],
    mutationFn: async () => {
      if (todo.complete) return alert("Todo is already completed");
      try {
        const res = await axios.patch(
          `${BASE_URL}/updatetodos/${todo.id}`,
          {
            complete: true, // Đặt trạng thái hoàn thành
          },
          {
            headers: {
              "Content-Type": "application/json",
              Authorization: `Bearer ${localStorage.getItem("token")}`, // Nếu cần token
            },
          }
        );
        return res.data;
      } catch (error) {
        console.log(error);
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todo"] });
    },
  });

  // Mutation để xóa todo
  const { mutate: deleteTodo, isPending: isDeleting } = useMutation({
    mutationKey: ["deleteTodo"],
    mutationFn: async () => {
      try {
        const res = await axios.delete(
          `${BASE_URL}/deletetodos/${todo.id}`,
          {
            headers: {
              Authorization: `Bearer ${localStorage.getItem("token")}`, // Nếu cần token
            },
          }
        );
        return res.data;
      } catch (error) {
        console.log(error);
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todo"] });
    },
  });

  return (
    <Flex gap={2} alignItems={"center"}>
      <Flex
        flex={1}
        alignItems={"center"}
        border={"1px"}
        borderColor={"gray.600"}
        p={2}
        borderRadius={"lg"}
        justifyContent={"space-between"}
      >
        <Text
          color={todo.complete ? "green.200" : "yellow.100"}
          textDecoration={todo.complete ? "line-through" : "none"}
        >
          {todo.title}
        </Text>
        {todo.complete && (
          <Badge ml="1" colorScheme="green">
            Done
          </Badge>
        )}
        {!todo.complete && (
          <Badge ml="1" colorScheme="yellow">
            In Progress
          </Badge>
        )}
      </Flex>
      <Flex gap={2} alignItems={"center"}>
        <Box
          color={"green.500"}
          cursor={"pointer"}
          onClick={() => updateTodo()}
        >
          {!isUpdating && <FaCheckCircle size={20} />}
          {isUpdating && <Spinner size={"sm"} />}
        </Box>
        <Box
          color={"red.500"}
          cursor={"pointer"}
          onClick={() => deleteTodo()}
        >
          {!isDeleting && <MdDelete size={25} />}
          {isDeleting && <Spinner size={"sm"} />}
        </Box>
      </Flex>
    </Flex>
  );
};

export default TodoItem;
