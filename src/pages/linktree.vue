<script setup lang="ts">
const links = [
  {
    img: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAANsAAADmCAMAAABruQABAAAAilBMVEUAAAD////Kysrv7+/39/fl5eWXl5dbW1v7+/vW1tbr6+vy8vLExMT8/Px2dnbk5OTAwMCAgIDT09OioqK0tLS4uLimpqZMTEyQkJDe3t6tra0lJSWGhoYtLS3Nzc1lZWVFRUU3NzcQEBBTU1MYGBgfHx+CgoJJSUl4eHhra2s+Pj46OjorKysLCwvdixKvAAAKvklEQVR4nO2da5uiPAyGOQmCjiKIHERB0VHU+f9/74WZd3cV2yYFOXQun6+LTO8NhTRJU0kelLTJyDSVaSElnIwb3kx6yZCa6MMKzWnkBmcj3e2lB11ipcmd+2Aba6qp2JH+OTN2EqRzWPvvdMWmFsZZuvoi8S9HkOdR2bzm32yNbaxNQsVe6vE5ue5hALbttFpDeCmbOg+V5Ur/TLJ03RCnIq97tvHYKmdOUM6cw2tpKjJqmI6bbWKFjl281maGX32ttSv+VyaC7UObO1PPDRa5f+sSpqrPl7B9qKNw60VunBv+pk+cR204X5j3bKHtxovk1a+BV8qtx2bnfQ8co2zCz7bs9LXQRDYnm5b1PeI/up2MPHBXq+I1TPumzLjYNF4v6KW6+fkicL2pM9c+HoamTmPiDw4mBxvssL5Uh51vzAI3sp1wBE0f5US6g45m+2wf52tnnD/1yFZMa8y5KjMvhPtdLBzbvBWadZolC91dTsO51XCNuSTdP0KxzV6Fc0zzc+x69jacVGZOQ6kk0yXw/5g0rg9z2J1+Jo4ZWvVWIVgRp80WZFP4eDbGbBFE9tS0tKbhDB4RRwk5mFIAGic1zvqPcbqkqUgzCGO7seMNUkIkupz1yCuMM2n3WeORSxrnivULibRsmQ2H6E5zkhPvM76Q0tfz9Ul34+UT8ZVOdzBJVyM/+z3IJg13RvvgkC7ev/Tr9FKpKWG8NAeTxCZt1G5HzCOdNOCAeCmRDfFd7E8m4Q0hpSQHk8LGuXrvVB/ECMHy+UIam5R3P2i0iN7zs4NJZZM2qHVEP7KupBFXJxKdrU60szsRXcX48RoW24A/dLK8JT5rDxFMJttwPZRCY5L3/OBgstmk9YAnnUx08+8imACbJE17HDwgCxgxyDbMSaeNQkchh/j+rVlhNsnolaKUplojc+pFq2CRJxmcGruNvn+GYJPWo65h1Lm5LWsX4rNx2qxrxPM9LBtXEJ5f2kQtrbL8tsqrcmI6mo0jCI/QZGQ6ZV48WCT+7tZSLjlGs3Hn9f6pFbMgFKHZaGskolnCv2a5btpN8bM04blYJ33IC7OEW9uL9NIsp14T4hURA2N0XT5XW9NxzNIsevEOu24GnJMkRyd/h3rNKr711ltvDUf79eaU+wMuZuLQLTWSWaxHS1sJLetvwbYjHt1hfcnyONBXnrI1mSUb48HDbU5GvghWUVlLY6lceUCz77E/6uu4y/LP+Mcqc546NJLSPlH2G7+cK+5qaW/DkdWUpSpO57iZjtcsX5RWmRZzpf08M2fJBY/KOq3kHJRWceajlitOSHpdhdNhVxolKIyiFFPlYwC5Vkq8kkO+q8x7MApCakOyY/3tMa2rIduxx9IgUA3ZhpySa8h27Hv4TDVj495a0amasWGqTftTM7ZBT7c325ttcHqzvdmGpjfbm61TfYytcDu1yxzmLCOn30Vh09TR3JkuV25wzg3/Ws2licsWpdBARGVzEHnm4bCpoaN4bhDPfETl2xQxDvI6shO2yU/5lRuX1VePZeFgTAK1Z428jmyTzYyCz+TErigBC4ygfV3fIld1tse2xOVRINNjam0oT3ZrbNht/Qc2Ghw/TXTarpO22PC7VhdMthHjl5vE3bJ207TERqz6poi59ZxlN+g91BIbz/Z3IBCohttIP2eEO0KB+nbY+LIMuIq+8bhaj9wPG3GDDF3InEJ1O1g/bJwZy41IbCvO++CeymGwhcmJr6IPVSE8DLYfaXNTWaJ8JukiGtu3kK9MzL7IwbGxHIt7ITYxCcvm/2I2disEwdkksCxIYDZwf5bAbGBTSZHZoMEKzQZs1RWaDdhXJzYbeyktOBuza4PgbMwt1qKzsfrNCs92/sVskvOL2fYCs4Ghy5h2q+Gz5WAQjBZoHj6bAaZmbgKzOdBtKSE9EdjgJA85pCcE2wd03524bOQeffciNoARg43c2udepBYpgrCBfRtScdnkCLo1IaQnCpucQvd+DjQLwwZuTsvEZYOLYp66LIrDJoO1NdVAs0BsoOtVDTQLxCafodtXAs0iscEldo+jF4kNdr0eQ3pCscGu10NITyw2uGXW/VZqsdhg1+u+dlQwNjmF/sRds13R2MAQ311ITzQ2cj/ye8XUSwfPJoMt8/4+leKxge1h/lY0i8cmL6C/8qf6X0A22PUyhWWDd6WsxWWDXa9AXDbY9QqFZYNdr524bHIK/SVXXDY4u2oJywa7Xqm4bHDUKxKXDe7MpwnLBrteSXVHhThscLvx6jGLArFhNgSLyga7XgKzab+YDXa9BGaTiacF/xI2vr2cYrHBrpfAbDLP+R2isfE0RRaNDU44Csz2gT/xQzg2DtdLPDa86yUgG/qMHgHZ0K4XNxvXmUrtsMlpS2zVBWAfbMg9Bdxs+IYx7bEhXS9uNi6Pri02nOvFzcbTDac9NpTrxc2G6vXVOhsc9arDxhW1aI8Nc8gePxtYZNUNG8L14mfjOJawTTbE81ODjaPZT5tssOtVgw3cUdING9wsqgYbRyitVTbZb4EN3xWtDtvFwR5cArletdjQQYua+013sY06BxzwkWqxybTTol/E9q08CsEzdNiuVz02OIuJZmMuEH2d3S6OHYutyYY87ghiG28DwOt92rzxKOaRUjXZkF3fmGyjyIDvAB0FztrZUpcN19OUzqYFX/DPSRuKqmKsJ7l7a/79F8wxx1Q2DRmaQHRnzCg/jVntUAE2TE6FykYbUEWoM2iIT3aA+EjS2WSvPluIQ0MeTf80504e6hgsBpv8WZsN425/xehzusJ70/ku+DAi2GQw5kVjY7zd9rssyWfxku9stcl0Fc/yWbAMOU6aY7KBywwa219n6bDenZJzELiRvXVCq9sT1Zhs4GmMNLZ55E3vD4ntR2w2yFsV7LyOyr+zX+Zis42Z7sWw2Z6+0NULmGu5YbM9La+ermB5zcNme3oTPl/C8JqHzfZkFcI19DbQw2Z7ckQJ19BjMoNme17xk66ies2DZnueS8TLaFmVIbMR3A7yhVfR2FRSOoFyqThs2kR1KJEayi/IKSOoa3AXGluh6diRqwcz47RjnjREuwVxSYY70+CF+iisMnem3nIVLGZ5tkEFnP7XF71vJ6ne9tQ2S2kUc+sVVolzI73e8IVszzrS2TTS/xEqrI9VaZSwNIpbGsXnKRPFyGf0WyV5zeBRWUwWdR6a0+VK1xd5ll7XPA9YDZ1ZvWRJXjP+JOixNhmZiu2t9PicZ7smT1c9rVhsxBQ0tcvnZBSGyv8zxb9sumepymSykbtiLryRVuj72MpyqpyLqdLy41VHe5nNhkzvDFIBwIaJNQ9VY4iN41y3gal46UFsNK956Co/ViBbsyLgvvTdSRxk499eOAD9+PQwG+5UzgHp8Oc8KAQb5y61fnWN/3XZxrARveYB6ZZmxizWV0vnsVIAw8a12akDHS/GbBHoka04oTqmF+Sg2LjP5Xul9ptTlsxid7W0nXBuceQacWyIdPHLdFinxnkRuMvCLPPJuEFKD8nWnn9yvJ6yfKG7kafwWeV1bOhyNlDrNCnN4tmKWawn2sy0otlqus37zcU3ZkHxEpua4VwFK/FeKTybrKK2T37tsvwc65E33YaqpnUKUxEHmyyPn2KcpVmSH7NszfC106WxuNjkMp5TLrcLlS+xVmdLc/0H6W3bNU7xLfUAAAAASUVORK5CYII=',
    title: 'Notion',
    link: 'https://marcustut.notion.site/ecba028a1c3448a18e44c6a9e0cd3c78',
  },
  {
    img: 'https://firebasestorage.googleapis.com/v0/b/mentu-lxs.appspot.com/o/Icon.png?alt=media&token=d72bc58e-b45a-401d-b370-9b961213352c',
    title: 'Student Portal',
    link: 'https://mentu-lxs.netlify.app',
    disabled: true,
    onclick: () => alert('Coming Soon!😎'),
  },
  {
    img: 'data:image/png;base64,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',
    title: 'Producer Video',
    link: 'https://mentu-lxs.netlify.app/video/pdvideo',
  },
  {
    img: 'https://i.pinimg.com/favicons/f446bd086a965df9beea6b3962eaa174079c624972e899cf3a39331e.png?d23e369be383e6616adf07b90a351957',
    title: 'Zoom',
    link: 'https://zoom.us/j/6905199217?pwd=bmdkTkVSUlZuNjNLTnFMVW5RVnQzQT09',
  },
]
</script>

<template>
  <div w="screen" h="screen" bg="textDark">
    <div
      overflow="y-hidden"
      display="flex"
      flex="col"
      justify="center"
      align="items-center"
      p="x-4"
      text="light-50"
      w="max-2xl"
      m="x-auto"
      h="full"
    >
      <img
        src="https://firebasestorage.googleapis.com/v0/b/mentu-lxs.appspot.com/o/Icon.png?alt=media&token=d72bc58e-b45a-401d-b370-9b961213352c"
        w="20"
        h="20"
      />
      <h3 font="bold" m="t-4">
        门徒练习生
      </h3>
      <a
        v-for="(link) in links"
        :key="link.title"
        :href="!link.disabled ? link.link : undefined"
        target="_blank"
        class="text-textDark hover:text-white"
        transition="duration-200 ease-in-out"
        bg="white hover:textDark"
        border="2 white dashed"
        w="full"
        m="t-6"
        p="x-4 y-6"
        display="flex"
        justify="center"
        align="items-center"
        pos="relative"
        cursor="pointer"
        :onclick="link.onclick ? link.onclick : undefined"
      >
        <img :src="link.img" pos="absolute left-4" w="12" h="12" />
        <p m="l-4" font="medium">
          {{ link.title }}
        </p>
      </a>
    </div>
  </div>
</template>
