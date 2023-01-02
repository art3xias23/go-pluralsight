#Aliases
#Git
alias ga='git add'
alias gcp='git cherry-pick'
alias gaa='git add -A'
alias gcm='git commit -m'
alias gpull='git pull'
alias gpush='git push'
alias gfo='git fetch origin'
alias gs='git status'
alias gl='git log'
alias gwl='git worktree list'
#MSBuild
alias msbuild='msbuild.exe'
#Nuget
alias nuget='nuget.exe'
#Random
alias python39='C:\\Users\\kmilchev\\AppData\\Local\\Programs\\Python\\Python39\\python.exe'
alias refreshbash='. ~kmilchev/.bashrc'
#Python
alias mssql-cli='python39 -m mssqlcli.main'
alias goog='python39 C:\\ProgramData\\chocolatey\\lib\\googler\\tools\\googler\\googler.py'
alias googler='LC_ALL="en_GB.UTF-8" goog'
#NPM
alias nis='npm install --save'
alias ni='npm install'
alias ns='npm start'

#Functions

##SqlCmd
function tsql
{
    # -u                       unicode output
    # -y 0                     don't limit field width
    # -h -1                    don't print headers
    # -s $'\t'                 use tab as field separator
    # -S $1                    first arg is the server name
    # -d $2                    second arg is the server name
    # -E                       use trusted connection
    # -Q "SET NOCOUNT ON; $3"  don't print (n rows affected) third arg is query 
    
    sqlcmd -u -y 0 -Y 30 -s $'\t' -S "$1" -d "$2" -E -Q "SET NOCOUNT ON; $3";
}

#Save Command History
PROMPT_COMMAND='history -a'

##MsSql-Cli

#Interactive
function msql
{
 
    # -E                        use integrated authentication
    # -S $1						first arg Server 
	# -d $2 					second args Database
	# --row-limit				
	# -Q "SET NOCOUNT ON; $3"	Query, don't print rows affected, third arg is query
    
    mssql-cli -E -S "$1" -d "$2" ;
}

function msqli
{
 
    # -E                        use integrated authentication
    # -S $1						first arg Server 
	# -d $2 					second args Database
	# --row-limit				
	# -Q "SET NOCOUNT ON; $3"	Query, don't print rows affected, third arg is query
    
    mssql-cli -E -S "$1" -d "$2" -i "$3" ;
}

function CheckRPAAP
{
 

    
    mssql-cli -E -S "apcore" -d "Release" -i ~kmilchev/Desktop/RPA_AP.sql;
}

function CheckRPAAPP
{
 

    
    mssql-cli -E -S "apcore" -d "APP_Release" -i ~kmilchev/Desktop/RPA_APP.sql;
}

#Refresh bashrc

[ -f "/c/Users/kmilchev/bin/win-sudo/s/path.sh" ] && source "/c/Users/kmilchev/bin/win-sudo/s/path.sh"
