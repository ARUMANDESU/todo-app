import './index.css';
import React from 'react'
import {createRoot} from 'react-dom/client'
import App from './App'
import {Toaster} from "@/components/ui/sonner";
import {ThemeProvider} from "@/components/theme-provider";

const container = document.getElementById('root')

const root = createRoot(container!)

root.render(
    <React.StrictMode>
        <ThemeProvider defaultTheme="system">
            <App/>
            <Toaster />
        </ThemeProvider>
    </React.StrictMode>,
)
