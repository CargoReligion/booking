// src/lib/api.ts
import axios from 'axios';
import type { User, SlotData, SlotDetails, CreateSessionFeedback, SessionFeedback, CreateSlotData, ApiResponse, Paginated } from '../types';
import { browser } from '$app/environment';

let initialUserId: string | null = null;

if (browser) {
  const storedUser = localStorage.getItem('currentUser');
  initialUserId = storedUser ? JSON.parse(storedUser).id : null;
}
const API_BASE_URL = import.meta.env.VITE_API_URL;

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
    return request;
  });

export const api = {
  createSlot: (slotData: CreateSlotData): Promise<ApiResponse<SlotData>> => 
    axiosInstance.post(`/api/slots`, slotData),

  getUpcomingSlots: (page: number = 1, pageSize: number = 10) => {
    return axiosInstance.get<Paginated<SlotData>>('/api/slots/upcoming', {
      params: { page, pageSize }
    })
    .then(response => response.data)
    .catch(error => {
      console.error('Error in getUpcomingSlots:', error);
      throw error;
    });
  },

  getAvailableSlots: (coachId: string, page: number = 1, pageSize: number = 10) => {
    return axiosInstance.get<Paginated<SlotData>>(`/api/slots/available/${coachId}`, {
      params: { page, pageSize }
    })
    .then(response => response.data)
    .catch(error => {
      console.error('Error in getAvailableSlots:', error);
      throw error;
    });
  },

  bookSlot: (id: string) => 
    axiosInstance.post<SlotData>(`/api/slots/${id}/book`),

  getUpcomingBookingsForStudent: (page: number = 1, pageSize: number = 10) => {
    return axiosInstance.get<Paginated<SlotData>>('/api/students/bookings', {
        params: { page, pageSize }
    })
      .then(response => response.data)
      .catch(error => {
        console.error('Error in getUpcomingBookingsForStudent:', error);
        throw error;
      });
  },

  getSlotDetails: (slotId: string) => {
    return axiosInstance.get<SlotDetails>(`/api/slots/${slotId}/details`)
      .then(response => response.data)
      .catch(error => {
        console.error('Error in getSlotDetails:', error);
        throw error;
      });
  },

  createSessionFeedback: (feedbackData: CreateSessionFeedback) => 
    axiosInstance.post<ApiResponse<SessionFeedback>>(`/api/session-feedback`, feedbackData),

  getStudentsWithSessions: () => {
    return axiosInstance.get<User[]>('/api/session-feedback/studentswithsessions')
      .then(response => response.data)
      .catch(error => {
        console.error('Error in getUsersWithSessions:', error);
        throw error;
      });
  },

  getSessionFeedbackForStudent: (studentId: string) => {
    return axiosInstance.get<SessionFeedback[]>(`/api/session-feedback/sessionsforstudent/${studentId}`)
      .then(response => response.data)
      .catch(error => {
        console.error('Error in getSessionFeedbackForStudent:', error);
        throw error;
      });
  },

  getAllUsers: () => 
    axiosInstance.get<User[]>(`/api/users`, {
      headers: {
        'X-User-Id': '989f159e-4ad7-4589-8a1c-1276078022ec',
      },
    }).then(response => response.data),
};

export const setUserID = (userId: string) => {
    axiosInstance.defaults.headers['X-User-Id'] = userId;
  };