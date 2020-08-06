#include "fqlib.c"
#include <stdio.h>



int main(int argc, char *argv[]){
	QUEUE* q =NULL;

qmanage(&q, 1, 10);

for (int i=0;i<10;i++){

put_on_queue(q, i);

}

q->size = 8;

qmanage(&q, 0, 10);

int y;

for (int j=0;j<10;j++){

take_off_queue(q, &y);

printf("%d ",y);}
}
