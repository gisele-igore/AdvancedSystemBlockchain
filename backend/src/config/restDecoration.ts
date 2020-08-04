import { ResponseObject } from '@loopback/rest';

export function response200(objectName: Object, objectType: string, examples: any) {
    const RESPONSE_200: ResponseObject = {
        description: `${objectType} Successful operation`,
        content: {
            'application/json': {
                'Content-Type': { 'x-ts-type': objectName },
                example: examples,
            },
        },
    }
    return RESPONSE_200
}
export function response200_1(objectName: Object, objectType: string, examples: Object) {
    const RESPONSE_200: ResponseObject = {
        description: `${objectType} Successful operation`,
        content: {
            'application/json': {
                'Content-Type': { 'x-ts-type': objectName },
                example: examples,
            },
        },
    }
    return RESPONSE_200
}

export function response204(objectType: Object) {
    const RESPONSE_204: ResponseObject = {
        description: 'No content',
        content: { 'application/json': { 'x-ts-type': objectType } },
    }
    return RESPONSE_204
}

export function response409(objectType: Object) {
    const RESPONSE_409: ResponseObject = {
        description: 'Resource Conflict',
        content: { 'application/json': { 'x-ts-type': objectType } },
    }
    return RESPONSE_409
}

export function response404(objectType: String) {
    const RESPONSE_404: ResponseObject = {
        description: `${objectType} not found`,
        content: { 'application/json': { 'x-ts-type': objectType } },
    }
    return RESPONSE_404
}

export function jsonFormatter(returnObject: Buffer) {
    const ts = JSON.parse(JSON.stringify(returnObject.toString()))
    let response = JSON.parse(`{${ts}}`)
    return response
}
