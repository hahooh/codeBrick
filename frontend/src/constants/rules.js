export function required(input) {
    return !!input || 'Requried';
}

export function minLength(input, minLength) {
    return !!input && input.length >= minLength || `Input minimum length is ${minLength}`;
}

export function maxLength(input, maxLength) {
    return !!input && input.length <= maxLength || `Input maximum length is ${maxLength}`;
}

export function onlyInteger(input) {
    return /^\d+$/.test(input) || 'Only integers';
}

export function onlyFloat(input) {
    return /^\d+\.\d+$/.test(input) || 'Only floats';
}

export function integerRange(input, from, to) {
    let targeInt = null;
    if (onlyInteger(input)) {
        targeInt = parseInt(input);
    } else if (onlyFloat(input)) {
        targeInt = parseFloat(input);
    }
    return targeInt !== null && from <= targeInt && targeInt <= to || `Input range from ${from} to ${to}`;
}