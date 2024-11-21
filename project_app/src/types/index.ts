export interface DataRecord {
    id: number
    Age : number
    Education : string
    Marital_Status :  string
    Occupation : string
    Income : string
}

export interface DataPagination {
    page: number
    page_size: number
    total_pages: number
    total_records: number
}

