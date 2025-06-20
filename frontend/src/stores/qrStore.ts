import { writable } from 'svelte/store';

export const detectedQrCode = writable<string | null>(null);
export const qrResult = writable<string | null>(null);
export const qrImage = writable<string | null>(null);
export const isScanning = writable<boolean>(false);
