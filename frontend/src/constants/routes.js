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
                required: true,
                defaultValue: '',
            },
            {
                text: 'Model',
                value: 'ModelName',
                input: true,
                type: INPUT_FIELD_TEXT,
                rules: [required],
                required: true,
                defaultValue: ''
            },
            {
                text: 'Make',
                value: 'Producer',
                input: true,
                type: INPUT_FIELD_TEXT,
                rules: [required],
                required: true,
                defaultValue: ''
            },
            {
                text: 'Year',
                value: 'Year',
                input: true,
                type: INPUT_FIELD_TEXT,
                rules: [required, input => integerRange(input, 1900, 2020)],
                required: true,
                typeCorrection: input => parseInt(input),
                defaultValue: ''
            },
            {
                text: 'MSRP',
                value: 'MSRP',
                input: true,
                type: INPUT_FIELD_TEXT,
                rules: [required, onlyFloat],
                required: true,
                typeCorrection: input => parseFloat(input),
                defaultValue: ''
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
                required: true,
                defaultValue: ''
            },
            {
                text: 'Booked',
                value: 'Booked',
                input: true,
                type: INPUT_FIELD_CHECKBOX,
                required: true,
                defaultValue: false,
            },
            {
                text: 'Listed',
                value: 'Listed',
                input: true,
                type: INPUT_FIELD_CHECKBOX,
                required: true,
                defaultValue: false,
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