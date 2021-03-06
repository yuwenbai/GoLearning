import { observable, computed, action, autorun } from 'mobx'
import { NavigationActions } from 'react-navigation'
import utils from '../utils'
import { DeviceEventEmitter } from 'react-native'

class account {
  @observable badge: number = 6

  @action.bound add () {
    let newBadge = this.badge
    newBadge++
    this.badge = newBadge
    // utils.global.navigator.dispatch(NavigationActions.setParams({
    //   params: { badge: newBadge },
    //   key: 'Account'
    // }))
    utils.global.navigator.dispatch(NavigationActions.navigate({
      routeName: 'Extra'
    }))
    
  }
  @action.bound OnBack () {
    // utils.global.navigator.dispatch(NavigationActions.setParams({
    //   params: { badge: newBadge },
    //   key: 'Account'
    // }))
    utils.global.navigator.dispatch(NavigationActions.back())
    
  }
}

export default account
