// src/lib/api.ts
import axios from 'axios';
import type { User, SlotData, SessionFeedback, CreateSlotData, ApiResponse, Paginated } from '../types';
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

  getUpcomingSlots: (page: number = 1, pageSize: number = 10) => {
    return axiosInstance.get<Paginated<SlotData>>('/slots/upcoming', {
      params: { page, pageSize }
    })
    .then(response => response.data)
    .catch(error => {
      console.error('Error in getUpcomingSlots:', error);
      throw error;
    });
  },

  getAvailableSlots: (coachId: string, page: number = 1, pageSize: number = 10) => {
    return axiosInstance.get<Paginated<SlotData>>(`/slots/available/${coachId}`, {
      params: { page, pageSize }
    })
    .then(response => response.data)
    .catch(error => {
      console.error('Error in getAvailableSlots:', error);
      throw error;
    });
  },

  bookSlot: (id: string) => 
    axiosInstance.post<ApiResponse<SlotData>>(`${API_BASE_URL}/slots/${id}/book`),

  getUpcomingBookingsForStudent: () => {
    return axiosInstance.get<Paginated<SlotData>>('/students/bookings')
      .then(response => response.data)
      .catch(error => {
        console.error('Error in getUpcomingBookingsForStudent:', error);
        throw error;
      });
  },

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