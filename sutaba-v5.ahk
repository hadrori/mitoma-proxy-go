#Requires AutoHotkey v2.0

global resultId := ""
global winId1 := ""
global winId2 := ""
global resultDelayMs := 80
global reloadDelayMs := 90
global sutabaDelayMs := 25
global counter := 0


F13:: TriggerNextBattle()
;; F14:: bookmark to quest (shortkeys)
F15:: PrepareResultWindow()

F16:: RegisterWindow1()
F17:: RegisterWindow2()
F18:: RegisterResultWindow()

F19:: DecreaseResultDelay()
F20:: IncreaseResultDelay()
^F19:: DecreaseReloadDelay()
^F20:: IncreaseReloadDelay()
+F19:: DecreaseSutabaDelay()
+F20:: IncreaseSutabaDelay()

F22:: LoadBattle()

DecreaseResultDelay() {
    global resultDelayMs
    if (resultDelayMs > 5) {
        resultDelayMs -= 5
    }
    ShowDelay(resultDelayMs, "resultDelayMs")
}

IncreaseResultDelay() {
    global resultDelayMs
    resultDelayMs += 5
    ShowDelay(resultDelayMs, "resultDelayMs")
}

DecreaseReloadDelay() {
    global reloadDelayMs
    if (reloadDelayMs > 5) {
        reloadDelayMs -= 5
    }
    ShowDelay(reloadDelayMs, "reloadDelayMs")
}

IncreaseReloadDelay() {
    global reloadDelayMs
    reloadDelayMs += 5
    ShowDelay(reloadDelayMs, "reloadDelayMs")
}

DecreaseSutabaDelay() {
    global sutabaDelayMs
    if (sutabaDelayMs > 5) {
        sutabaDelayMs -= 5
    }
    ShowDelay(sutabaDelayMs, "sutabaDelayMs")
}

IncreaseSutabaDelay() {
    global sutabaDelayMs
    sutabaDelayMs += 5
    ShowDelay(sutabaDelayMs, "sutabaDelayMs")
}

ShowDelay(delayMs, name) {
    Toast(name ": " delayMs " ms")
}

Toast(msg) {
    Tooltip(msg)
    SetTimer(() => ToolTip(), -1000)
}

RegisterWindow1() {
    global winId1
    winId1 := WinGetID("A")
    Toast("1st window regietered: " winId1)
}

RegisterWindow2() {
    global winId2
    winId2 := WinGetID("A")
    Toast("2nd window regietered: " winId2)
}

RegisterResultWindow() {
    global resultId
    resultId := WinGetID("A")
    Toast("result window regietered: " resultId)
}

; Prepare result window
PrepareResultWindow() {
    url := ""
    A_Clipboard := ""
    SendInput("^l")
    Sleep(5)
    SendInput("^c")
    Sleep(5)
    SendInput("{Escape}")
    Sleep(5)

    url := ClipWait(0.05) ? A_Clipboard : ""

    A_Clipboard := url ? StrReplace(url, "raid", "result") : "https://game.granbluefantasy.jp/#quest/index"

    WinActivate(resultId)
    Sleep(2)
    SendInput("^l")
    Sleep(2)
    SendInput("^v")
}

ReloadResult(resultId) {
    WinActivate(resultId)
    Sleep(2)
    SendInput("{Enter}")
    A_Clipboard := ""
}

FriendAlert() {
    global counter

    static threashold := 35
    if (++counter >= threashold) {
        counter -= threashold
        SoundBeep()
    }
}

LoadSupporters() {
    global reloadDelayMs

    Sleep(reloadDelayMs)
    SendInput("{F5}")
}

LoadBattle() {
    global winId2, sutabaDelayMs

    ; SendInput("{F5}")
    ; Sleep(2)

    WinActivate(winId2)
    Sleep(sutabaDelayMs)
    SendInput("{F14}")
}

TriggerNextBattle() {
    global resultId, resultDelayMs

    ; reload result
    ReloadResult(resultId)

    ; wait before trigger the next battle
    Sleep(resultDelayMs)
    Click()
    Sleep(1)

    ; starburst
    LoadSupporters()

    ;FriendAlert()
}
