import { required, minLength, maxLength, integerRange, onlyFloat } from "./rules";

export const subtitles = [
    {
        title: 'Inventory',
        url: '/inventory',
        method: 'fetchInventories',
        headers: [
            { text: 'NO', value: 'Id', input: false },
            {
                text: 'Vin#',
                value: 'VehicleIdentificationNumber',
                input: true,
                type: 'text',
                rules: [required, input => minLength(input, 17), input => maxLength(input, 17)],
                required: true
            },
            {
                text: 'Model',
                value: 'ModelName',
                input: true,
                type: 'text',
                rules: [required],
                required: true
            },
            {
                text: 'Make',
                value: 'Producer',
                input: true,
                type: 'text',
                rules: [required],
                required: true
            },
            {
                text: 'Year',
                value: 'Year',
                input: true,
                type: 'text',
                rules: [required, input => integerRange(input, 1900, 2020)],
                required: true
            },
            {
                text: 'MSRP',
                value: 'MSRP',
                input: true,
                type: 'text',
                rules: [required, onlyFloat],
                required: true
            },
            {
                text: 'Status',
                value: 'Status',
                input: true,
                type: 'select',
                items: [
                    'ordered',
                    'in stock',
                    'sold'
                ],
            },
            { text: 'Booked', value: 'Booked', input: true, type: 'checkbox' },
            { text: 'Listed', value: 'Listed', input: true, type: 'checkbox' },
        ]
    },
    {
        title: 'Commission',
        url: '/commision',
        method: 'fetchCommistions',
    },
    {
        title: 'Manage Market',
        url: '/market',
        method: 'fetchMarkets'
    },
    {
        title: 'Manage Customer',
        url: '/user',
        method: 'fetchUsers',
    },
    {
        title: 'Report Setting',
        url: '/setting',
        method: 'fetchSettings',
    },
    {
        title: 'Sign Out',
        url: '/auth',
        method: 'logOut'
    }
];