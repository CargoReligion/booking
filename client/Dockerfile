# Use Node.js as the base image
FROM node:18

# Set the working directory inside the container
WORKDIR /app

# Copy package files first (to leverage Docker's caching)
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the app's code
COPY . .

# Expose the app's port
EXPOSE 3000

# Run the SvelteKit development server
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]
