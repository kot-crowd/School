strComputer = "."
strOwner = "A111111"
strProcess = "'notepad.exe'"
' Connect to WMI service and Win32_Process filtering by name'
Set objWMIService = GetObject("winmgmts:{impersonationLevel=impersonate}!\\" _
    & strComputer & "\root\cimv2")
Set colProcessbyName = objWMIService.ExecQuery("Select * from Win32_Process Where Name = " _
    & strProcess)
' Get the process ID for the process started by the user in question'
For Each objProcess in colProcessbyName
    colProperties = objProcess.GetOwner(strUsername,strUserDomain)
    if strUsername = strOwner then
        strProcessID = objProcess.ProcessId
    end if
next
' We have the process ID for the app in question for the user, now we kill it'
Set colProcessList = objWMIService.ExecQuery("Select * from Win32_Process where ProcessId =" & strProcessID)
For Each objProcess in colProcess
    objProcess.Terminate()
Next