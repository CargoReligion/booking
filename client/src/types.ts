 export interface User {
    id: string;
    name: string;
    phoneNumber: string;
    role: string;
  }
  
  export interface SlotData {
    id: string;
    coachId: string;
    coachName: string;
    startTime: string;
    endTime: string;
    studentId?: string;
    booked: boolean;
  }

  export interface SlotDetails {
    id: string;
    coachId: string;
    coachName: string;
    startTime: string;
    endTime: string;
    studentId?: string;
    booked: boolean;
    coachPhoneNumber: string;
    studentPhoneNumber: string;
    studentName: string;
  }
  
  export interface CreateSlotData {
    startTime: string;
  }
  
  export interface CreateSessionFeedback {
    slotId: string;
    satisfaction: number;
    notes: string;
  }

  export interface SessionFeedback {
    id: string;
    slotId: string;
    satisfaction: number;
    notes: string;
    createdAt: string;
  }

  export interface Paginated<T> {
    data: T[];
    page: number;
    pageSize: number;
    totalCount: number;
    totalPages: number;
}
  
  export interface ApiResponse<T> {
    data: T;
    // Add other properties that your API returns, if any
  }