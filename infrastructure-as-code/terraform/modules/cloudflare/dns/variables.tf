variable "cloudflare_api_token" {
  description = "cloudflare api token"
  type        = string
}

variable "cloudflare_zone_id" {
  description = "cloudflare zone id"
  type        = string
}

variable "A_records" {
  description = "DNS A records to add. It is a map of IP => { value: domain name }"
  type        = map(object({
    value = string
    ttl   = optional(number, 120),
  }))
}

variable "TXT_records" {
  description = "DNS TXT records to add, It is map of key => { value: answer }"
  type        = map(object({
    value = string
    ttl   = optional(number, 120),
  }))
}

variable "CNAME_records" {
  description = "DNS CNAME records to add, it is a map of domain => { value: cname }"
  type        = map(object({
    value = string
    ttl   = optional(number, 120),
  }))
}

variable "use_cloudflare_proxy" {
  description = "should we use cloudflare proxy for provided domain"
  type        = bool
  default     = false
g