import React from 'react'
import { View, Text, StyleSheet, TouchableOpacity } from 'react-native'
import Splash from 'react-native-splash-screen'
import { inject, observer } from 'mobx-react'

@inject('account')
@observer
export default class LoginScreen extends React.Component {
  componentDidMount () {
    setTimeout(() => Splash.hide(), 100)
  }
  OnClick_SureWithdraw(){
    this.props.account.OnBack()
  }
  render () {
    console.log(this.props)
    return (
      <View style={styles.header}>
        <TouchableOpacity style={styles.loginBtn} onPress={() => this.OnClick_SureWithdraw()} activeOpacity={0.2} focusedOpacity={0.5}>
        <Text>请登录!</Text>
        </TouchableOpacity>
      </View>
    )
  }
}
const styles = StyleSheet.create({
  container: {
      height: "100%",
      width: "100%",
      backgroundColor: "#F2F2F2",
  },
  header: {
      height: 44,
      width: "100%",
      backgroundColor: "#fff",
      display: "flex",
      flexDirection: "row",
      alignItems: "center",
      justifyContent: "space-between",
      paddingLeft: 10,
      paddingRight: 10,
  },
  headerText: {
      textAlign: "center",
      color: "black",
      fontSize: 16,
  },
  tipsText: {
      color: "#008AFF",
      fontSize: 13,
      marginLeft: 5,
  },
  SettingMenu: {
      flex: 1,
      height: 500,
      backgroundColor: "#fff",
      borderRadius: 10,
      margin: 10,
      padding: 20,
  },
  withdrawstateStyle: {
      marginTop: 15,
      marginLeft: 10,
      alignItems: "stretch",
      flexDirection: "row",
  },
  title_text: {
      marginLeft: 10,
      color: "#000000",
      fontSize: 18,
  },
  detail_text: {
      color: "#999999",
      fontSize: 16,
  },
  selectText: {
      color: "#00BD9A",
      fontSize: 16,
  },
  detail_text_time: {
      color: "#999999",
      fontSize: 14,
  },
  selectText_time: {
      color: "#00BD9A",
      fontSize: 14,
  },
  selectText_failed: {
      color: "#E75A5A",
      fontSize: 16,
  },
  selectText_failed_time: {
      color: "#E75A5A",
      fontSize: 14,
  },
  loginBtnBg: {
      justifyContent: "center",
      alignItems: "center",
      width: "100%",
      marginTop: 40,
  },
  loginBtn: {
      justifyContent: 'center',
      alignItems: 'center',
      width: "90%",
      height: 40,
      backgroundColor: '#E75A5A',
      borderRadius: 20,
  },
  loginBtn_no: {
      justifyContent: 'center',
      alignItems: 'center',
      width: "90%",
      height: 40,
      backgroundColor: '#CCCCCC',
      borderRadius: 20,
  }
})