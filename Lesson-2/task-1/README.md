## Задание 1

Выполните сборку ваших предыдущих программ под операционную систему,
отличающуюся от текущей. Проанализируйте вывод команды file для полученного
исполняемого файла. Попробуйте запустить исполняемый файл.

### file numericPipelineDarwin
numericPipelineDarwin: Mach-O 64-bit x86_64 executable

./numericPipelineDarwin     
zsh: Ошибка формата выполняемого файла: ./numericPipelineDarwin

### file numericPipelineWindows.exe 
numericPipelineWindows.exe: PE32+ executable (console) x86-64 (stripped to external PDB), for MS Windows

./numericPipelineWindows.exe 
zsh: Ошибка формата выполняемого файла: ./numericPipelineWindows.exe

### file numericPipelineLinux      
numericPipelineLinux: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, Go BuildID=qSWSwppF5RQWbeo8cn2v/deAStNIr-y1CNMKq1y8H/YwkRr-u_aKJb-EKz7Vcx/RXXeKD-Av4MpDNeUSXFE, not stripped

./numericPipelineLinux 
2
4
6
8
10
12
14
16
18
20
22
24
