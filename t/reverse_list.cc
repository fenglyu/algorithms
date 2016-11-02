//创建:用q指向当前链表的最后一个节点；用p指向即将插入的新节点。
//反向:用p和q反转工，r记录链表中剩下的还未反转的部分。

//#include "stdafx.h"
#include <iostream>
using namespace std;

struct ActList
{
    char ActName[20];
    char Director[20];
    int Mtime;
    ActList *next;
};

ActList* head;

ActList*  Create()
{//start of CREATE()
    ActList* p=NULL;
    ActList* q=NULL;
    head=NULL;
    int Time;
    cout<<"Please input the length of the movie."<<endl;
    cin>>Time;
    while(Time!=0){
    p=new ActList;
    //类似表达:  TreeNode* node = new TreeNode;//Noice that [new] should be written out.
    p->Mtime=Time;
    cout<<"Please input the name of the movie."<<endl;
    cin>>p->ActName;
    cout<<"Please input the Director of the movie."<<endl;
    cin>>p->Director;

    if(head==NULL)
    {
    head=p;
    }
    else
    {
    q->next=p;
    }
    q=p;
    cout<<"Please input the length of the movie."<<endl;
    cin>>Time;
    }
    if(head!=NULL)
    q->next=NULL;
    return head;

}//end of CREATE()


void DisplayList(ActList* head)
{//start of display
    cout<<"show the list of programs."<<endl;
    while(head!=NULL)
    {
        cout<<head->Mtime<<"\t"<<head->ActName<<"\t"<<head->Director<<"\t"<<endl;
        head=head->next;
    }
}//end of display


ActList* ReverseList2(ActList* head)
{
    //ActList* temp=new ActList;
 if(NULL==head|| NULL==head->next) return head;
    ActList* p;
    ActList* q;
    ActList* r;
    p = head;
    q = head->next;
    head->next = NULL;
    while(q){
        r = q->next; //
        q->next = p;
        p = q; //
        q = r; //
    }
    head=p;
    return head;
}

ActList* ReverseList3(ActList* head)
{
    ActList* p;
    ActList* q;
    p=head->next;
    while(p->next!=NULL){
        q=p->next;
        p->next=q->next;
        q->next=head->next;
        head->next=q;
    }

    p->next=head;//相当于成环
    head=p->next->next;//新head变为原head的next
    p->next->next=NULL;//断掉环
    return head;
}
int main(int argc, char* argv[])
{
//  DisplayList(Create());
//  DisplayList(ReverseList2(Create()));
    DisplayList(ReverseList3(Create()));
    return 0;
}
