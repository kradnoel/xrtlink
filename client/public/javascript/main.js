new Vue({
  el: '#app',
  template: `
  <div class='row'>
    <div class='col-md-12'>
      <input v-model='link' type='text' placeholder='Enter the link to shrink...' />
      <button @click='clicked(link)'>Shrink</button>
    </div>
    <div v-if="shrinked != '' " class='col-md-12'>
      <div class="container3">
        <span style='font-family: arial; font-weight:200; color: rgba(3, 3, 3, .80); padding: 5px;' v-if='shrinked'>
          <small v-if='showLink === true'>Shrinked link:</small>
          <small style='font-family: arial; font-weight:500; color: rgba(0, 0, 0, .80);'>{{shrinked}}</small>
        </span>
      </div>
    </div>
  </div>
  `,
  data : function() {
    return {
      link: '',
      shrinked: '',
      showLink: false
    }
  },
  methods: {
    clicked: function(e) {
      if(this.link === '') {
        this.showLink = false
        this.shrinked = 'Link cannot be empty!!! Try again'
        //console.log('value is empty')
      } else {
        axios({
          method: 'post',
          url: 'http://localhost:8080',
          headers: {"Content-Type": "application/json", "Accept": "application/json"},
          data: {
            Link: this.link
          }
        }).then((r) => {
          //console.log(r.config)
          var data = r.data

          if(data.error === true) {
            this.showLink = false
            this.shrinked = r.data.data
          }else {
            this.showLink = true
            this.shrinked = 'http://localhost:3000/' + r.data.data
          }

          //console.log(r.data.data)
          this.link = ''
        }).catch((e) => {
          //console.log(e.config)
        })
      }
    }
  }
})
