interface Marker {
  iconPath: string
  id: number
  latitude: number
  longitude: number
  width: number
  height: number
}

const defaultAvatar = '/resources/car.png'
const initialLat = 29.761267625855936
const initialLng = 121.87264654736123

Page({
  isPageShowing: false,
  socket: undefined as WechatMiniprogram.SocketTask | undefined,

  data: {
    avatarURL: '',
    setting: {
      skew: 0,
      rotate: 0,
      showLocation: true,
      showScale: true,
      subKey: '',
      layerStyle: -1,
      enableZoom: true,
      enableScroll: true,
      enableRotate: false,
      showCompass: false,
      enable3D: false,
      enableOverlooking: false,
      enableSatellite: false,
      enableTraffic: false,
    },
    location: {
      latitude: initialLat,
      longitude: initialLng,
    },
    scale: 16,
    markers: [] as Marker[],
  },
  onMyLocationTap() {
    wx.getLocation({
      type: 'gcj02',
      success: res => {
        this.setData({
          location: {
            latitude: res.latitude,
            longitude: res.longitude,
          },
        })
      },
      fail: () => {
        wx.showToast({
          icon: 'none',
          title: '请前往设置页授权',
        })
      }
    })
  },
  moveCars() {
    const map = wx.createMapContext("map")
    const dest = {
      latitude: initialLat,
      longitude: initialLng,
    }
    // map.translateMarker({
    //   destination: {
    //         latitude: newLat,
    //         longitude: newLng,
    //       },
    // })
  },
  onHide() {
    this.isPageShowing = false;
    if (this.socket) {
      this.socket.close({
        success: () => {
          this.socket = undefined
        }
      })
    }
  },
})
