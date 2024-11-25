// src/lib/api.ts
import axios from 'axios';
import type { User, SlotData, SessionFeedback, CreateSlotData, ApiResponse } from '../types';
import { browser } from '$app/environment';

let initialUserId: string | null = null;

if (browser) {
  const storedUser = localStorage.getItem('currentUser');
  initialUserId = storedUser ? JSON.parse(storedUser).id : null;
}
const API_BASE_URL = 'http://localhost:8080/api';

const axiosInstance = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

if (initialUserId) {
  axiosInstance.defaults.headers['X-User-Id'] = initialUserId;
}

// Add an interceptor to log headers for each request
axiosInstance.interceptors.request.use(request => {
    console.log('Request headers:', request.headers);
    return request;
  });

export const api = {
  createSlot: (slotData: CreateSlotData): Promise<ApiResponse<SlotData>> => 
    axiosInstance.post(`${API_BASE_URL}/slots`, slotData),

  getUpcomingSlots: () => 
    axiosInstance.get<SlotData[]>(`${API_BASE_URL}/slots/upcoming`).then(response => response.data),

  getAvailableSlots: () => 
    axiosInstance.get<ApiResponse<SlotData[]>>(`${API_BASE_URL}/slots/available`),

  bookSlot: (id: string) => 
    axiosInstance.post<ApiResponse<SlotData>>(`${API_BASE_URL}/slots/${id}/book`),

  getUpcomingBookingsForStudent: () => 
    axiosInstance.get<ApiResponse<SlotData[]>>(`${API_BASE_URL}/students/bookings`),

  getSlotDetails: (id: number) => 
    axiosInstance.get<ApiResponse<SlotData>>(`${API_BASE_URL}/slots/${id}/details`),

  createSessionFeedback: (feedbackData: SessionFeedback) => 
    axiosInstance.post<ApiResponse<SessionFeedback>>(`${API_BASE_URL}/session-feedback`, feedbackData),

  getPastSessionFeedbacks: () => 
    axiosInstance.get<ApiResponse<SessionFeedback[]>>(`${API_BASE_URL}/session-feedback/past`),

  getAllUsers: () => 
    axiosInstance.get<User[]>(`${API_BASE_URL}/users`, {
      headers: {
        'X-User-Id': '989f159e-4ad7-4589-8a1c-1276078022ec',
      },
    }).then(response => response.data),
};

export const setUserID = (userId: string) => {
    axiosInstance.defaults.headers['X-User-Id'] = userId;
  };