export interface DataRecord {
    id: number
    Age : number
    Work_Class : string
    Education : string
    Marital_Status :  string
    Occupation : string
    Relationship: string
    Race: string
    Sex: string
    Capital_Gain: number
    Capital_Loss: number
    Hours_Per_Week: number
    Native_Country: number
    Income : string
}

export interface DataPagination {
    page: number
    page_size: number
    total_pages: number
    total_records: number
}

export interface User {
    id: number
    username: string
    password: string
}

export interface JWTPayload {
    MapClaims: {
        eat: number
        iat: number
    }
    Session: string
}

export interface UserFilters{
    education: string | undefined
    marital_status: string | undefined
    occupation: string | undefined
    income: string | undefined
    order_by: string | undefined
    order_direction: string | undefined
    min_age: string | undefined
    max_age: string | undefined
}

