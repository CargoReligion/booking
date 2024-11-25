// src/types.ts
export interface User {
    id: string;
    name: string;
    phoneNumber: string;
    role: string;
  }
  
  export interface SlotData {
    id: string;
    coachId: string;
    startTime: string;
    endTime: string;
    studentId?: string;
    booked: boolean;
  }
  
  export interface CreateSlotData {
    startTime: string;
  }
  
  export interface SessionFeedback {
    id: string;
    slotId: string;
    satisfaction: number;
    notes: string;
  }
  
  export interface ApiResponse<T> {
    data: T;
    // Add other properties that your API returns, if any
  }