export function formatDate(date: Date, timeZone: string = 'America/New_York'): string {
    return date.toLocaleString('en-US', {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
        hour12: true,
        timeZone: timeZone
    });
}


export function localToUTC(date: Date): Date {
    return new Date(date.getTime() - (date.getTimezoneOffset() * 60000));
}

export function UTCToLocal(dateString: string): Date {
    const date = new Date(dateString);
    return new Date(date.getTime() + (date.getTimezoneOffset() * 60000));
}

export function UTCToEST(dateString: string): Date {
    const date = new Date(dateString);
    return new Date(date.toLocaleString('en-US', {timeZone: 'America/New_York'}));
}