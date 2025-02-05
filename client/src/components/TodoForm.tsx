import { Button, Flex, Input, Spinner } from "@chakra-ui/react";
import { useState } from "react";
import { IoMdAdd } from "react-icons/io";
import axios from "axios";
import { BASE_URL } from "../App"; // Đảm bảo BASE_URL đúng
import { useMutation, useQueryClient } from "@tanstack/react-query";

const TodoForm = () => {
  const [newTodo, setNewTodo] = useState("");
  const queryClient = useQueryClient();

  const { mutate: createTodo, isPending: isCreating } = useMutation({
    mutationKey: ["createTodo"],
    mutationFn: async (e: React.FormEvent) => {
      e.preventDefault();

      if (!newTodo.trim()) {
        alert("Please enter a valid todo!");
        return;
      }

      try {
        // Gửi yêu cầu POST đến API
        const response = await axios.post(
          `${BASE_URL}/createtodos`,
          {
            body: newTodo,
          },
          {
            headers: {
              Authorization: `Bearer ${localStorage.getItem("token")}`, // Lấy accessToken từ localStorage
            },
          }
        );

        // Kiểm tra kết quả
        if (response.status === 201) {
          setNewTodo(""); // Reset input
          queryClient.invalidateQueries({ queryKey: ["todo"] }); // Cập nhật lại dữ liệu todo
        } else {
          alert("Failed to create todo");
        }
      } catch (error: any) {
        alert(error.message); // Xử lý lỗi
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todo"] });
    },
  });

  return (
    <form onSubmit={createTodo}>
      <Flex gap={2}>
        <Input
          type="text"
          value={newTodo}
          onChange={(e) => setNewTodo(e.target.value)}
          ref={(input) => input && input.focus()}
        />
        <Button
          mx={2}
          type="submit"
          _active={{
            transform: "scale(.97)",
          }}
          isLoading={isCreating}
          loadingText="Adding"
        >
          {!isCreating && <IoMdAdd size={30} />}
          {isCreating && <Spinner size="xs" />}
        </Button>
      </Flex>
    </form>
  );
};

export default TodoForm;
