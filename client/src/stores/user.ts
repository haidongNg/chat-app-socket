import type { User } from "../models/user.model";
import { writable } from "svelte/store";

export const user = writable({} as User);
export const isAuthenticated = writable(false);