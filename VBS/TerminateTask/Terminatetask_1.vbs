Sub KillProc( myProcess )
'Authors: Denis St-Pierre and Rob van der Woude
'Purpose: Kills a process and waits until it is truly dead

    Dim blnRunning, colProcesses, objProcess
    blnRunning = False

    Set colProcesses = GetObject( _
                       "winmgmts:{impersonationLevel=impersonate}" _
                       ).ExecQuery( "Select * From Win32_Process", , 48 )
    For Each objProcess in colProcesses
        If LCase( myProcess ) = LCase( objProcess.Name ) Then
            ' Confirm that the process was actually running
            blnRunning = True
            ' Get exact case for the actual process name
            myProcess  = objProcess.Name
            ' Kill all instances of the process
            objProcess.Terminate()
        End If
    Next

    If blnRunning Then
        ' Wait and make sure the process is terminated.
        ' Routine written by Denis St-Pierre.
        Do Until Not blnRunning
            Set colProcesses = GetObject( _
                               "winmgmts:{impersonationLevel=impersonate}" _
                               ).ExecQuery( "Select * From Win32_Process Where Name = '" _
                             & myProcess & "'" )
            WScript.Sleep 100 'Wait for 100 MilliSeconds
            If colProcesses.Count = 0 Then 'If no more processes are running, exit loop
                blnRunning = False
            End If
        Loop
        ' Display a message
        WScript.Echo myProcess & " was terminated"
    Else
        WScript.Echo "Process """ & myProcess & """ not found"
    End If
End Sub