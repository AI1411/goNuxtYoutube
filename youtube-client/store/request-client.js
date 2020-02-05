export class RequestClient {
  constructor(axios) {
    this.axios = axios
  }

  async get(uri, params = {}) {
    const querystring = Object.keys(params).map(key => key + '=' + params[key]).join('&');
    const query = querystring.length > 0 ? `${uri}?${querystring}` : uri
    return await this.axios.$get(query)
  }
}

export function createRequestClient(axios) {
  return new RequestClient(axios)
}
