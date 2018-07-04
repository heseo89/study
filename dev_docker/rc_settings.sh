#!/bin/bash

# bashrc
BASHRC_STR=`cat <<EOF
parse_git_branch() {
    git branch 2> /dev/null | sed -e '/^[^*]/d' -e 's/* \(.*\)/ (\1)/'          
}                                                                               
export PS1="\e[0;34m\u \[\033[32m\]\w\[\033[1;33m\]\n\$(parse_git_branch)\[\033[00m\] $ "
EOF`
echo "$BASHRC_STR" >> ~/.bashrc

#  vimrc
VIMRC_STR=`cat <<EOF
" Syntax Highlighting                                  
if has("syntax")                                       
    syntax on                                          
endif                                                  
                                                           
syntax on                                              
filetype indent plugin on                              
                                                       
let python_highlight_all = 1                           
                                                       
set nowrap                                             
set tabstop=4                                          
set expandtab                                          
set cindent                                            
" set smarttab                                         
set shiftwidth=4                                       
set softtabstop=4                                      
set bg=dark                                            
set history=1000                                       
set textwidth=100                                      
set mouse-=a                                           
set nu                                                 
set nocindent                                          
set nobackup                                           
                                                       
set hlsearch                                           
set autoindent                                         
set ruler                                              
                                                       
set cc=80                                              
" highlight ColorColumn ctermbg=72 guibg=#5faf87       
                                                       
map<F4> :w<Enter>:!python %<Enter>                     
map<F9> :w<Enter>:!nosetests --nocapture %<Enter>      
map<F8> :w<Enter>:!sh %<Enter>                         
map<F12> :w<Enter>:!/usr/local/go/bin/go run %<Enter>  
EOF`
echo "$VIMRC_STR" >> ~/.vimrc
