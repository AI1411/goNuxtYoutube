import {createRequestClient} from "./request-client";

export const state = () => ({
  items: [],
  meta: {}
});

export const actions = {
  async fetchPopularVideos({commit}, payload) {
    //APIにリクエストするためのクライアントを作成
    const client = createRequestClient(this.$axios);
    //GETリクエストを送信する
    const res = await client.get(payload.uri, payload.params);

    //resをコミットに渡す
    commit('mutatePopularVideos', res)
  }
};

export const mutations = {
  mutatePopularVideos(state, payload) {
    //stateにAPIのレスポンスをセットする
    state.items = payload.items ? state.items.concat(payload.items) : [];
    state.meta = payload;
  },
};

//componentからstateを参照するため
export const getters = {
  getPopularVideos(state) {
    return state.items
  },
  getMeta(state) {
    return state.meta
  }
};
