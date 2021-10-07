use platform

fn homedir {
  if (==s $platform:os "windows") {
    put $E:USERPROFILE
  }
  put $E:HOME
}

fn editor {

  if (has-env VISUAL) {
    put $E:VISUAL
    return
  } elif (has-env EDITOR) {
    put $E:EDITOR
    return
  } elif (==s $platform:os "windows") {
    put 'notepad.exe'
    return
  } elif (==s $platform:os "darwin") {
    put 'TextPad'
    return
  }

}
