import { QueryClient, useMutation } from "@tanstack/react-query";
import toast from "react-hot-toast";

export function useEditCabin() {
  const { mutate: editCabin, isLoading: isEditing } = useMutation({
    mutationFn: ({ newCabinData, id }) => editCabin(newCabinData, id),
    onSuccess: () => {
      toast.success("cabin successfull edited");
      QueryClient.invalidateQueries({ queryKey: ["cabins"] });
      //   reset();
    },
    onError: (err) => toast.error(err.message),
  });
  return { isEditing, editCabin };
}
