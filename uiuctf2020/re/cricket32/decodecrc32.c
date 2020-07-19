#include "stdio.h"
#include "stdbool.h"

bool crc_check(int constant, int seg) {
	int result;

	__asm__(
		"crc32 %%ebx, %%eax\n"
		:"=a"(result)
		:"b"(seg), "a"(constant)
	);


	if ((result ^ seg) == 0) {
		return true;
	}

	return false;
}

int main() {
	int constants[7] = {0x4c10e5f7, 0xf357d2d6, 0x373724ff, 0x90bbb46c, 0x34d98bf6, 0xd79ee67d, 0xaa79007c};
	int flag[7] = {0x63756975, 0x617b6674, 0x5f5f5f5f, 0x5f5f5f5f, 0x5f5f5f5f, 0x5f5f5f5f, 0x007d5f5f};

	
	for (int i = 0; i < 7; ++i) {

		while (!crc_check(constants[i], flag[i])) {
			flag[i]++;
			
			if ((flag[i] & 0x000000ff) == 0x0000007e) {
				flag[i] &= 0xffffff5f;
				flag[i] += 0x100;
				if ((flag[i] & 0x0000ff00) == 0x7e00) {
					flag[i] &= 0xffff5fff;
					flag[i] += 0x10000;
					if ((flag[i] & 0x00ff0000) == 0x7e0000) {
						flag[i] &= 0xff5fffff;
						flag[i] += 0x1000000;
					}
				}
			}
			printf("%08x\r", flag[i]);
		}
		printf("%08x\n", flag[i]);
	}

	return 0;
}
