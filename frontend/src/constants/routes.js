import { required, minLength, maxLength, integerRange, onlyFloat } from "./rules";
import { INPUT_FIELD_TEXT, INPUT_FIELD_CHECKBOX, INPUT_FIELD_SELECT } from './inputFieldType'

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
                type: INPUT_FIELD_TEXT,
                rules: [
                    required,
                    input => minLength(input, 17),
                    input => maxLength(input, 17)
                ],
                required: true
            },
            {
                text: 'Model',
                value: 'ModelName',
                input: true,
                type: INPUT_FIELD_TEXT,
                rules: [required],
                required: true
            },
            {
                text: 'Make',
                value: 'Producer',
                input: true,
                type: INPUT_FIELD_TEXT,
                rules: [required],
                required: true
            },
            {
                text: 'Year',
                value: 'Year',
                input: true,
                type: INPUT_FIELD_TEXT,
                rules: [required, input => integerRange(input, 1900, 2020)],
                required: true
            },
            {
                text: 'MSRP',
                value: 'MSRP',
                input: true,
                type: INPUT_FIELD_TEXT,
                rules: [required, onlyFloat],
                required: true
            },
            {
                text: 'Status',
                value: 'Status',
                input: true,
                type: INPUT_FIELD_SELECT,
                items: [
                    'ordered',
                    'in stock',
                    'sold'
                ],
                required: true
            },
            {
                text: 'Booked',
                value: 'Booked',
                input: true,
                type: INPUT_FIELD_CHECKBOX,
                required: true
            },
            {
                text: 'Listed',
                value: 'Listed',
                input: true,
                type: INPUT_FIELD_CHECKBOX,
                required: true
            },
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